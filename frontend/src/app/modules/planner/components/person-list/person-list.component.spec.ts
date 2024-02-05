import { ComponentFixture, DeferBlockBehavior, DeferBlockState, TestBed } from '@angular/core/testing';

import { PersonListComponent } from './person-list.component';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { TestbedHarnessEnvironment } from '@angular/cdk/testing/testbed';
import { HarnessLoader } from '@angular/cdk/testing';
import { MatProgressBarHarness } from '@angular/material/progress-bar/testing';
import { PersonWithMetadata } from '@app/shared/interfaces/person';
import { MatCardHarness } from '@angular/material/card/testing';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { MatTooltipHarness } from '@angular/material/tooltip/testing';
import { DragDropModule } from '@angular/cdk/drag-drop';

const mockPersons: PersonWithMetadata[] = [
  {
    first_name: 'John',
    last_name: 'Doe',
    working_hours: 8,
    id: '1',
  } as PersonWithMetadata,
];

describe('PersonListComponent', () => {
  let component: PersonListComponent;
  let fixture: ComponentFixture<PersonListComponent>;
  let mockPersonDataContainerService: jasmine.SpyObj<PersonDataContainerService>;
  let loader: HarnessLoader;

  beforeEach(async () => {
    mockPersonDataContainerService = jasmine.createSpyObj('PersonDataContainerService', ['persons', 'persons$'], {
      persons$: mockPersons,
      persons: mockPersons,
    });

    await TestBed.configureTestingModule({
      deferBlockBehavior: DeferBlockBehavior.Manual,
      imports: [PersonListComponent, DragDropModule],
      providers: [
        provideNoopAnimations(),
        {
          provide: PersonDataContainerService,
          useValue: mockPersonDataContainerService,
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    loader = TestbedHarnessEnvironment.loader(fixture);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should display placeholder text in defer block', async () => {
    // should have it by default
    expect(await loader.getHarness(MatProgressBarHarness)).toBeTruthy();
  });

  // KEEP THIS COMMENTED OUT

  it('should display persons', async () => {
    const deferBlock = (await fixture.getDeferBlocks())[0];

    await deferBlock.render(DeferBlockState.Complete);
    // get MatCards
    const matCards = await loader.getAllHarnesses(MatCardHarness);

    expect(matCards.length).toBe(1);
    expect(await matCards[0].getTitleText()).toBe(''); // no title
    expect(await matCards[0].getText()).toContain(`${mockPersons[0].id} (${mockPersons[0].working_hours})`);

    // get MatTooltips
    const matTooltips = await loader.getAllHarnesses(MatTooltipHarness);

    expect(matTooltips.length).toBe(1);

    await (await matTooltips[0].host()).hover();
    expect(await matTooltips[0].getTooltipText()).toBe(`${mockPersons[0].last_name}, ${mockPersons[0].first_name}`);
  });
});
