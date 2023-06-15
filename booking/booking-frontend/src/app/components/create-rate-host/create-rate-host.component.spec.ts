import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateRateHostComponent } from './create-rate-host.component';

describe('CreateRateHostComponent', () => {
  let component: CreateRateHostComponent;
  let fixture: ComponentFixture<CreateRateHostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateRateHostComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateRateHostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
