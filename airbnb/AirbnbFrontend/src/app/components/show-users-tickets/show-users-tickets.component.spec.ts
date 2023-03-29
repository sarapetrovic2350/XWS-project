import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShowUsersTicketsComponent } from './show-users-tickets.component';

describe('ShowUsersTicketsComponent', () => {
  let component: ShowUsersTicketsComponent;
  let fixture: ComponentFixture<ShowUsersTicketsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ShowUsersTicketsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ShowUsersTicketsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
