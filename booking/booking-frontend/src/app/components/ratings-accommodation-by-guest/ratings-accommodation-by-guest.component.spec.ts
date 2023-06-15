import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RatingsAccommodationByGuestComponent } from './ratings-accommodation-by-guest.component';

describe('RatingsAccommodationByGuestComponent', () => {
  let component: RatingsAccommodationByGuestComponent;
  let fixture: ComponentFixture<RatingsAccommodationByGuestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RatingsAccommodationByGuestComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RatingsAccommodationByGuestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
