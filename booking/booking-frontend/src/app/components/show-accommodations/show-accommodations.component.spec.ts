import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowAccommodationsComponent } from './show-accommodations.component';

describe('ShowAccommodationsComponent', () => {
  let component: ShowAccommodationsComponent;
  let fixture: ComponentFixture<ShowAccommodationsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ShowAccommodationsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ShowAccommodationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
