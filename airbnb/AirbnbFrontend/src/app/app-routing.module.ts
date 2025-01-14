import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {UserRegistrationComponent} from "./components/user-registration/user-registration.component";
import {LoginComponent} from "./components/login/login.component";
import {CreateFlightComponent} from "./components/create-flight/create-flight.component";
import {ShowFlightsComponent} from "./components/show-flights/show-flights.component";
import {HomeComponent} from "./components/home/home.component";
import { CreateTicketComponent } from './components/create-ticket/create-ticket.component';
import { ShowUsersTicketsComponent } from './components/show-users-tickets/show-users-tickets.component';

const routes: Routes = [
  {
    path: 'register-user',
    component: UserRegistrationComponent
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'createFlight',
    component: CreateFlightComponent
  },
  {
    path: 'showFlights',
    component: ShowFlightsComponent
  },
  {
    path: '',
    component: HomeComponent
  }, 
  {
    path: 'createTicket/:id', 
    component: CreateTicketComponent
  }, 
  {
    path: 'showUserTickets', 
    component: ShowUsersTicketsComponent
  }
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
