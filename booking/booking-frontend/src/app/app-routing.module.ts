import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {UserRegistrationComponent} from "./components/user-registration/user-registration.component";
import {LoginComponent} from "./components/login/login.component";
import { CreateAccommodationComponent } from './components/create-accommodation/create-accommodation.component';
import {ShowFlightsComponent} from "./components/show-flights/show-flights.component";
import {HomeComponent} from "./components/home/home.component";
import { ViewHostAccommodationComponent } from './components/view-host-accommodation/view-host-accommodation.component';
import { CreateAvailabilityComponent } from './components/create-availability/create-availability.component';
import { UpdateUserComponent } from './components/update-user/update-user.component';
import { ViewReservationsComponent } from './components/view-reservations/view-reservations.component';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  },
  {
    path: 'register-user',
    component: UserRegistrationComponent
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'create-accommodation',
    component: CreateAccommodationComponent
  },
  {
    path: 'show-host-accommodations',
    component: ViewHostAccommodationComponent
  },
  {
    path: 'createAvailability/:id', 
    component: CreateAvailabilityComponent
  },
  {
    path: 'update-user', 
    component: UpdateUserComponent
  }, 
  {
    path: 'view-reservations', 
    component: ViewReservationsComponent
  }

];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
