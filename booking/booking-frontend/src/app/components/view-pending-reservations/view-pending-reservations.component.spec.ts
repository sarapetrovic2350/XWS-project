import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewPendingReservationsComponent } from './view-pending-reservations.component';

describe('ViewPendingReservationsComponent', () => {
  let component: ViewPendingReservationsComponent;
  let fixture: ComponentFixture<ViewPendingReservationsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewPendingReservationsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewPendingReservationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
