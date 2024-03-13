import { DepartmentWithMetadata } from '@app/shared/interfaces/department';
import { DynamicDataSource, DynamicDatabase, DynamicFlatNode } from './tree-data-source';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';
import { FlatTreeControl } from '@angular/cdk/tree';
import { TestBed } from '@angular/core/testing';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { TimeslotAPIService } from '@app/shared/services/timeslot-api.service';
import { WorkplaceAPIService } from '@app/shared/services/workplace-api.service';
import { of } from 'rxjs';

describe('DynamicFlatNode', () => {
  const mockDepartment: DepartmentWithMetadata = {
    id: 'department1',
    name: 'department1',
    created_at: new Date(),
    updated_at: new Date(),
    deleted_at: null,
  };
  let dynamicFlatNode: DynamicFlatNode;

  beforeEach(() => {
    dynamicFlatNode = new DynamicFlatNode(mockDepartment, 0, 'department');
  });

  it('should create', () => {
    expect(dynamicFlatNode).toBeTruthy();
  });

  it('should have expandable property', () => {
    expect(dynamicFlatNode.expandable).toBe(true);

    dynamicFlatNode.type = 'workplace';
    expect(dynamicFlatNode.expandable).toBe(true);

    dynamicFlatNode.type = 'timeslot';
    expect(dynamicFlatNode.expandable).toBe(false);
  });

  it('should have isLoading property', () => {
    expect(dynamicFlatNode.isLoading).toBe(false);
  });
});

describe('DynamicDatabase', () => {
  const mockDepartments: DepartmentWithMetadata[] = [
    {
      id: 'department1',
      name: 'department1',
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
    },
    {
      id: 'department2',
      name: 'department2',
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
    },
  ];
  let dynamicDatabase: DynamicDatabase;

  beforeEach(() => {
    dynamicDatabase = new DynamicDatabase(mockDepartments);
  });

  it('should create', () => {
    expect(dynamicDatabase).toBeTruthy();
  });

  it('should handle null departments', () => {
    dynamicDatabase = new DynamicDatabase(null);
    expect(dynamicDatabase).toBeTruthy();
  });

  it('should have getChildren method', () => {
    const mockNode: WorkplaceWithMetadata = {
      id: 'workplace1',
      name: 'workplace1',
      department_id: 'department1',
      created_at: new Date(),
      updated_at: new Date(),
      deleted_at: null,
    };
    const children = dynamicDatabase.getChildren(mockNode);
    expect(children).toBeUndefined();

    dynamicDatabase.setData(mockDepartments[0], [mockNode]);
    const children2 = dynamicDatabase.getChildren(mockDepartments[0]);
    expect(children2).toEqual([mockNode]);
  });
});

describe('DynamicDataSource', () => {
  let mockWorkplaceAPIService: jasmine.SpyObj<WorkplaceAPIService>;
  let mockTimeslotAPIService: jasmine.SpyObj<TimeslotAPIService>;
  let mockDepartmentAPIService: jasmine.SpyObj<DepartmentAPIService>;

  let mockTreeControl: jasmine.SpyObj<FlatTreeControl<DynamicFlatNode>>;
  let dynamicDataSource: DynamicDataSource;

  beforeEach(() => {
    mockWorkplaceAPIService = jasmine.createSpyObj('WorkplaceAPIService', ['getWorkplaces']);
    mockTimeslotAPIService = jasmine.createSpyObj('TimeslotAPIService', ['getTimeslots']);
    mockDepartmentAPIService = jasmine.createSpyObj('DepartmentAPIService', ['getDepartments']);
    mockDepartmentAPIService.getDepartments.and.returnValue(of({ data: [], status: 200, message: 'success' }));

    mockTreeControl = jasmine.createSpyObj('FlatTreeControl', ['dataNodes']);

    TestBed.configureTestingModule({
      providers: [
        { provide: WorkplaceAPIService, useValue: mockWorkplaceAPIService },
        { provide: TimeslotAPIService, useValue: mockTimeslotAPIService },
        { provide: DepartmentAPIService, useValue: mockDepartmentAPIService },
      ],
    });

    TestBed.runInInjectionContext(() => {
      dynamicDataSource = new DynamicDataSource(mockTreeControl);
    });
  });

  it('should create', () => {
    expect(dynamicDataSource).toBeTruthy();
  });

  it('should have dataChange property', () => {
    const mockNode = new DynamicFlatNode(
      {
        id: 'department1',
        name: 'department1',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      0,
      'department',
    );
    dynamicDataSource.data = [mockNode];

    expect(dynamicDataSource.data).toEqual([mockNode]);
    expect(dynamicDataSource.getDataChange().value).toEqual([mockNode]);
  });

  it('toggleNode should call getWorkplaces if type === department', () => {
    const mockNode: DynamicFlatNode = new DynamicFlatNode(
      {
        id: 'department1',
        name: 'department1',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      0,
      'department',
    );
    dynamicDataSource.toggleNode(mockNode, false);

    expect(mockWorkplaceAPIService.getWorkplaces).toHaveBeenCalled();
  });

  it('toggleNode should call getTimeslots if type === workplace', () => {
    const mockNode: DynamicFlatNode = new DynamicFlatNode(
      {
        id: 'workplace1',
        name: 'workplace1',
        department_id: 'department1',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      1,
      'workplace',
    );
    dynamicDataSource.toggleNode(mockNode, false);

    expect(mockTimeslotAPIService.getTimeslots).toHaveBeenCalled();
  });

  it('handleTreeControl should call toggleNode if added', () => {
    spyOn(dynamicDataSource, 'toggleNode');
    const mockNode: DynamicFlatNode = new DynamicFlatNode(
      {
        id: 'department1',
        name: 'department1',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      0,
      'department',
    );

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    dynamicDataSource.handleTreeControl({ added: [mockNode] } as any);

    expect(dynamicDataSource.toggleNode).toHaveBeenCalledTimes(1);
    expect(dynamicDataSource.toggleNode).toHaveBeenCalledWith(mockNode, true);
  });

  it('handleTreeControl should call toggleNode if removed', () => {
    spyOn(dynamicDataSource, 'toggleNode');
    const mockNode: DynamicFlatNode = new DynamicFlatNode(
      {
        id: 'department1',
        name: 'department1',
        created_at: new Date(),
        updated_at: new Date(),
        deleted_at: null,
      },
      0,
      'department',
    );

    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    dynamicDataSource.handleTreeControl({ removed: [mockNode, mockNode] } as any);

    expect(dynamicDataSource.toggleNode).toHaveBeenCalledTimes(2);
    expect(dynamicDataSource.toggleNode).toHaveBeenCalledWith(mockNode, false);
  });
});
