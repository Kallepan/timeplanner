import { ComponentFixture, TestBed, fakeAsync, tick } from '@angular/core/testing';

import { PersonAutocompleteComponent } from './person-autocomplete.component';
import { HarnessLoader } from '@angular/cdk/testing';
import { PersonDataContainerService } from '@app/shared/services/person-data-container.service';
import { provideNoopAnimations } from '@angular/platform-browser/animations';
import { MatAutocompleteHarness } from '@angular/material/autocomplete/testing';

describe('PersonAutocompleteComponent', () => {
  let component: PersonAutocompleteComponent;
  let fixture: ComponentFixture<PersonAutocompleteComponent>;
  let mockPersonDataContainerService: jasmine.SpyObj<PersonDataContainerService>;
  let loader: HarnessLoader;

  beforeEach(async () => {
    mockPersonDataContainerService = jasmine.createSpyObj('PersonDataContainerService', [], {
      persons$: [
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        { last_name: 'Doe', id: '123', first_name: 'John' } as any,
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        { last_name: 'Smith', id: '456', first_name: 'Kalle' } as any,
      ],
    });

    await TestBed.configureTestingModule({
      imports: [PersonAutocompleteComponent],
      providers: [{ provide: PersonDataContainerService, useValue: mockPersonDataContainerService }, provideNoopAnimations()],
    }).compileComponents();

    fixture = TestBed.createComponent(PersonAutocompleteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should display the person last name and id', () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const person = { last_name: 'Doe', id: '123' } as any;
    expect(component.displayFn(person)).toBe('Doe (123)');
  });

  it('should subscribe to the persons observable', fakeAsync(() => {
    const oldSubscription = component.filteredPersons$.subscribe((persons) => {
      expect(persons.length).toBe(2);
    });
    tick(150);
    oldSubscription.unsubscribe();

    component.filteredPersons$.subscribe((persons) => {
      expect(persons.length).toBe(1);
    });
    // simulate a change in the input
    const input = fixture.nativeElement.querySelector('input');
    input.value = 'Doe';
    input.dispatchEvent(new Event('input'));
    fixture.detectChanges();
    tick(150);
  }));

  it('should correctly display the persons in the compoenent', async () => {
    component.filteredPersons$.subscribe((persons) => {
      expect(persons.length).toBe(1);
    });

    const matAutocomplete = await loader.getHarness(MatAutocompleteHarness);
    expect(matAutocomplete).toBeTruthy();

    await matAutocomplete.focus();
    await matAutocomplete.enterText('Doe');
    await (await matAutocomplete.host()).dispatchEvent('focusin');
    fixture.detectChanges();
    expect(await matAutocomplete.getValue()).toBe('Doe');

    expect(await matAutocomplete.isOpen()).toBe(true);
  });

  it('should _filter the persons correctly', () => {
    // filter by last name
    const persons = component.filter('Doe');
    expect(persons.length).toBe(1);
    expect(persons[0].last_name).toBe('Doe');

    // filter by id
    const persons2 = component.filter('123');
    expect(persons2.length).toBe(1);
    expect(persons2[0].id).toBe('123');

    // filter by first name
    const persons3 = component.filter('John');
    expect(persons3.length).toBe(1);
    expect(persons3[0].first_name).toBe('John');

    // filter by non existing person
    const persons4 = component.filter('NonExisting');
    expect(persons4.length).toBe(0);
  });
});
