import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewHostAccommodationComponent } from './view-host-accommodation.component';

describe('ViewHostAccommodationComponent', () => {
  let component: ViewHostAccommodationComponent;
  let fixture: ComponentFixture<ViewHostAccommodationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewHostAccommodationComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewHostAccommodationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
