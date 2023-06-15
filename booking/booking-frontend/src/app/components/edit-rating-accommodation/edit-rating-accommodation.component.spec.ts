import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditRatingAccommodationComponent } from './edit-rating-accommodation.component';

describe('EditRatingAccommodationComponent', () => {
  let component: EditRatingAccommodationComponent;
  let fixture: ComponentFixture<EditRatingAccommodationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EditRatingAccommodationComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EditRatingAccommodationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
