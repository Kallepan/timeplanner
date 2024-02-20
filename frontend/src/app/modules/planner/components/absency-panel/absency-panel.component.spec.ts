import { ComponentFixture, DeferBlockBehavior, DeferBlockState, TestBed } from '@angular/core/testing';

import { AbsencyPanelComponent } from './absency-panel.component';
import { MatBottomSheetRef } from '@angular/material/bottom-sheet';
import { AbsencyDataContainerService } from '../../services/absency-data-container.service';
import { ActiveWeekHandlerService } from '@app/shared/services/active-week-handler.service';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { ActiveDepartmentHandlerService } from '@app/shared/services/active-department-handler.service';
import { of } from 'rxjs';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { MatProgressBarHarness } from '@angular/material/progress-bar/testing';
import { HarnessLoader } from '@angular/cdk/testing';

describe('AbsencyPanelComponent', () => {
  let component: AbsencyPanelComponent;
  let fixture: ComponentFixture<AbsencyPanelComponent>;
  let mockBottomSheetRef: jasmine.SpyObj<MatBottomSheetRef<AbsencyPanelComponent>>;
  let loader: HarnessLoader;

  // for absencyDataContainer
  let mockDepartmentAPIService: jasmine.SpyObj<DepartmentAPIService>;
  let activeWeekHandlerService: ActiveWeekHandlerService;
  let mockActiveDepartmentHandlerService: jasmine.SpyObj<ActiveDepartmentHandlerService>;

  beforeEach(async () => {
    mockDepartmentAPIService = jasmine.createSpyObj('DepartmentAPIService', ['getAbsencesForDepartment']);
    mockDepartmentAPIService.getAbsencesForDepartment.and.returnValue(
      of({
        data: [],
        message: 'test',
        status: 200,
      }),
    );

    mockActiveDepartmentHandlerService = jasmine.createSpyObj('ActiveDepartmentHandlerService', [], {
      activeDepartment$: 'test',
    });

    mockBottomSheetRef = jasmine.createSpyObj('MatBottomSheetRef', ['dismiss']);

    await TestBed.configureTestingModule({
      imports: [AbsencyPanelComponent],
      providers: [
        AbsencyDataContainerService,
        ActiveWeekHandlerService,
        { provide: DepartmentAPIService, useValue: mockDepartmentAPIService },
        { provide: ActiveDepartmentHandlerService, useValue: mockActiveDepartmentHandlerService },
        { provide: MatBottomSheetRef, useValue: mockBottomSheetRef },
      ],
      deferBlockBehavior: DeferBlockBehavior.Manual,
    }).compileComponents();

    activeWeekHandlerService = TestBed.inject(ActiveWeekHandlerService);

    fixture = TestBed.createComponent(AbsencyPanelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    loader = TestbedHarnessEnvironment.loader(fixture);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should dismiss the bottom sheet', () => {
    component.close();
    expect(mockBottomSheetRef.dismiss).toHaveBeenCalled();
  });

  it('should have a default absencesGroupedByWeekday', () => {
    expect(component.absencyDataContainer.absencesGroupedByWeekday$.length).toEqual(7);
    expect(component.absencyDataContainer.loading$).toEqual(false);
  });

  it('should display progress-bar', async () => {
    const progressBar = await loader.getHarness(MatProgressBarHarness);
    expect(progressBar).toBeTruthy();
  });

  it('should not display progress-bar', async () => {
    const deferBlockFixture = (await fixture.getDeferBlocks())[0];
    await deferBlockFixture.render(DeferBlockState.Complete);

    expect(component.absencyDataContainer.loading$).toEqual(false);
    expect(fixture.nativeElement.querySelector('mat-progress-bar')).toBeFalsy();
  });

  it('mock an active week', async () => {
    mockDepartmentAPIService.getAbsencesForDepartment.and.returnValue(
      of({
        data: [
          {
            person_id: '1',
            date: '2021-01-01',
            reason: 'sick',
            created_at: new Date(),
          },
        ],
        message: 'test',
        status: 200,
      }),
    );

    activeWeekHandlerService.activeWeekByDate = new Date('2021-01-01');
    fixture.detectChanges();

    const deferBlockFixture = (await fixture.getDeferBlocks())[0];
    await deferBlockFixture.render(DeferBlockState.Complete);

    expect(mockDepartmentAPIService.getAbsencesForDepartment).toHaveBeenCalledWith('test', '2021-01-01');
    expect(component.absencyDataContainer.absencesGroupedByWeekday$.length).toEqual(7);
    expect(component.absencyDataContainer.absencesGroupedByWeekday$[4].absences.length).toEqual(1);

    // Check fixture html
    expect(fixture.nativeElement.innerHTML).toContain('1. 1 - sick');

    // fetch cells
    const cells = fixture.nativeElement.querySelectorAll('.cell');
    // 7 days and 2 cells per day since sick is returned by the mock for all days
    expect(cells.length).toEqual(14);

    // fetch content cells
    const contentCells = fixture.nativeElement.querySelectorAll('.content.cell');
    // 7 cells for each day
    expect(contentCells.length).toEqual(7);
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    contentCells.forEach((cell: any) => {
      expect(cell.innerHTML).toContain('1 - sick');
    });
  });
});
