import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateRateAccommodationComponent } from './create-rate-accommodation.component';

describe('CreateRateAccommodationComponent', () => {
  let component: CreateRateAccommodationComponent;
  let fixture: ComponentFixture<CreateRateAccommodationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateRateAccommodationComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateRateAccommodationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
