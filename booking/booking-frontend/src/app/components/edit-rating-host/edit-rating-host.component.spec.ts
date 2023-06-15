import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditRatingHostComponent } from './edit-rating-host.component';

describe('EditRatingHostComponent', () => {
  let component: EditRatingHostComponent;
  let fixture: ComponentFixture<EditRatingHostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EditRatingHostComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EditRatingHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
