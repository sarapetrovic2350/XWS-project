import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowRatingsForHostComponent } from './show-ratings-for-host.component';

describe('ShowRatingsForHostComponent', () => {
  let component: ShowRatingsForHostComponent;
  let fixture: ComponentFixture<ShowRatingsForHostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ShowRatingsForHostComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ShowRatingsForHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
