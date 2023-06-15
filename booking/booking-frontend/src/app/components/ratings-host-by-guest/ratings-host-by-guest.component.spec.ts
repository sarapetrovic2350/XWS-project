import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RatingsHostByGuestComponent } from './ratings-host-by-guest.component';

describe('RatingsHostByGuestComponent', () => {
  let component: RatingsHostByGuestComponent;
  let fixture: ComponentFixture<RatingsHostByGuestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RatingsHostByGuestComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RatingsHostByGuestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
