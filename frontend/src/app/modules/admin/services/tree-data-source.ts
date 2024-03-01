/*
  This service is used to create a tree structure of the departments, workplaces and timeslots
  It uses the Angular CDK Tree to create the tree structure
  It uses the DepartmentAPIService, WorkplaceAPIService and TimeslotAPIService to fetch the data
  It uses the DynamicFlatNode, DynamicDatabase and DynamicDataSource classes to create the tree structure
*/
import { CollectionViewer, DataSource, SelectionChange } from '@angular/cdk/collections';
import { FlatTreeControl } from '@angular/cdk/tree';
import { inject } from '@angular/core';
import { DepartmentWithMetadata } from '@app/shared/interfaces/department';
import { TimeslotWithMetadata } from '@app/shared/interfaces/timeslot';
import { WorkplaceWithMetadata } from '@app/shared/interfaces/workplace';
import { DepartmentAPIService } from '@app/shared/services/department-api.service';
import { TimeslotAPIService } from '@app/shared/services/timeslot-api.service';
import { WorkplaceAPIService } from '@app/shared/services/workplace-api.service';
import { BehaviorSubject, Observable, catchError, map, merge, of, switchMap, tap, throwError } from 'rxjs';

type ChildNode = WorkplaceWithMetadata | TimeslotWithMetadata;
type RootNode = DepartmentWithMetadata;
type LeafNode = TimeslotWithMetadata;

export class DynamicFlatNode {
  constructor(
    public item: RootNode | ChildNode,
    public level: number,
    public type: 'department' | 'workplace' | 'timeslot',
    public isLoading: boolean = false,
  ) {}

  get expandable(): boolean {
    return this.type !== 'timeslot';
  }
}

export class DynamicDatabase {
  private dataMap = new Map<RootNode | Exclude<ChildNode, LeafNode>, ChildNode[]>();

  departments: DynamicFlatNode[] = [];
  constructor(departments: RootNode[]) {
    this.departments = departments.map((department) => new DynamicFlatNode(department, 0, 'department'));
  }

  // Get All children of the node
  getChildren(node: RootNode | Exclude<ChildNode, LeafNode>): ChildNode[] | undefined {
    return this.dataMap.get(node);
  }

  // Add the node to the dataMap along with the children
  setData(node: RootNode | Exclude<ChildNode, LeafNode>, data: ChildNode[]): void {
    this.dataMap.set(node, data);
  }
}

export class DynamicDataSource implements DataSource<DynamicFlatNode> {
  private dataChange = new BehaviorSubject<DynamicFlatNode[]>([]);
  getDataChange() {
    // only for testing
    return this.dataChange;
  }
  get data(): DynamicFlatNode[] {
    return this.dataChange.value;
  }
  set data(value: DynamicFlatNode[]) {
    this._treeControl.dataNodes = value;
    this.dataChange.next(value);
  }

  private departmentAPIService = inject(DepartmentAPIService);
  private workplaceAPIService = inject(WorkplaceAPIService);
  private timeslotAPIService = inject(TimeslotAPIService);

  private _database: DynamicDatabase;
  constructor(private _treeControl: FlatTreeControl<DynamicFlatNode>) {
    this.departmentAPIService
      .getDepartments()
      .pipe(
        map((resp) => resp.data),
        catchError(() => of([])),
      )
      .subscribe((departments) => {
        this._database = new DynamicDatabase(departments);
        this.data = this._database.departments;
      });
  }
  connect(collectionViewer: CollectionViewer): Observable<DynamicFlatNode[]> {
    this._treeControl.expansionModel.changed.subscribe((change) => {
      if ((change as SelectionChange<DynamicFlatNode>).added || (change as SelectionChange<DynamicFlatNode>).removed) {
        this.handleTreeControl(change as SelectionChange<DynamicFlatNode>);
      }
    });

    return merge(collectionViewer.viewChange, this.dataChange).pipe(map(() => this.data));
  }

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  disconnect(_collectionViewer: CollectionViewer): void {}

  /** Handle expand/collapse behaviors */
  handleTreeControl(change: SelectionChange<DynamicFlatNode>) {
    if (change.added) {
      change.added.forEach((node) => this.toggleNode(node, true));
    }
    if (change.removed) {
      change.removed
        .slice()
        .reverse()
        .forEach((node) => this.toggleNode(node, false));
    }
  }

  /**
   * Toggle the node, remove from display list
   */
  toggleNode(node: DynamicFlatNode, expand: boolean) {
    if (node.type === 'timeslot') return;
    if (expand && node.isLoading) {
      return;
    }
    of(node)
      .pipe(
        tap(() => (node.isLoading = true)),
        switchMap((node) => {
          if (node.type === 'department') {
            const item = node.item as DepartmentWithMetadata;
            return this.workplaceAPIService.getWorkplaces(item.id).pipe(
              map((resp) => resp.data),
              catchError((err) => throwError(() => err)),
            );
          }

          if (node.type === 'workplace') {
            const item = node.item as WorkplaceWithMetadata;
            return this.timeslotAPIService.getTimeslots(item.department_id, item.id).pipe(
              map((resp) => resp.data),
              catchError((err) => throwError(() => err)),
            );
          }

          return of([]);
        }),
        tap(() => (node.isLoading = false)),
        map((children) => children.map((child) => new DynamicFlatNode(child, node.level + 1, node.type === 'department' ? 'workplace' : 'timeslot'))),
      )
      .subscribe({
        next: (children) => {
          const index = this.data.indexOf(node);
          if (index < 0) {
            // If no children, or cannot find the node, no op
            return;
          }

          if (expand) {
            this.data.splice(index + 1, 0, ...children);
          } else {
            let count = 0;
            for (let i = index + 1; i < this.data.length && this.data[i].level > node.level; i++, count++) {
              continue;
            }
            this.data.splice(index + 1, count);
          }

          // notify the change
          this.dataChange.next(this.data);
        },
        error: () => (node.isLoading = false),
      });
  }
}
