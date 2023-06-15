import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowRatingsAccommodationsHostComponent } from './show-ratings-accommodations-host.component';

describe('ShowRatingsAccommodationsHostComponent', () => {
  let component: ShowRatingsAccommodationsHostComponent;
  let fixture: ComponentFixture<ShowRatingsAccommodationsHostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ShowRatingsAccommodationsHostComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ShowRatingsAccommodationsHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
