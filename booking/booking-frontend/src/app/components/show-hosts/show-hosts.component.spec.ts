import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowHostsComponent } from './show-hosts.component';

describe('ShowHostsComponent', () => {
  let component: ShowHostsComponent;
  let fixture: ComponentFixture<ShowHostsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ShowHostsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ShowHostsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
