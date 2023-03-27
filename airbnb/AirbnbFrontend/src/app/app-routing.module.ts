import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {UserRegistrationComponent} from "./components/user-registration/user-registration.component";
import {LoginComponent} from "./components/login/login.component";
import {CreateFlightComponent} from "./components/create-flight/create-flight.component";
import {ShowFlightsComponent} from "./components/show-flights/show-flights.component";

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
  }
];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
