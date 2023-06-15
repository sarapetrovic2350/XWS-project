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
import { ViewPendingReservationsComponent } from './components/view-pending-reservations/view-pending-reservations.component';
import { ShowHostsComponent } from './components/show-hosts/show-hosts.component';
import { CreateRateHostComponent } from './components/create-rate-host/create-rate-host.component';
import { ShowAccommodationsComponent } from './components/show-accommodations/show-accommodations.component';
import { CreateRateAccommodationComponent } from './components/create-rate-accommodation/create-rate-accommodation.component';
import { RatingsHostByGuestComponent } from './components/ratings-host-by-guest/ratings-host-by-guest.component';
import { EditRatingHostComponent} from "./components/edit-rating-host/edit-rating-host.component";
import { ShowRatingsForHostComponent } from './components/show-ratings-for-host/show-ratings-for-host.component';
import {
  RatingsAccommodationByGuestComponent
} from "./components/ratings-accommodation-by-guest/ratings-accommodation-by-guest.component";
import {
  EditRatingAccommodationComponent
} from "./components/edit-rating-accommodation/edit-rating-accommodation.component";

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
  },
  {
    path: 'view-pending-reservations',
    component: ViewPendingReservationsComponent
  },
  {
    path: 'show-hosts',
    component: ShowHostsComponent
  },
  {
    path: 'rating-host/:email',
    component: CreateRateHostComponent
  },
  {
    path: 'show-accommodations',
    component: ShowAccommodationsComponent
  },

  {
    path: 'rating-accommodation/:id',
    component: CreateRateAccommodationComponent
  },
  {
    path: 'ratings-host-by-guest',
    component: RatingsHostByGuestComponent
  },
  {
    path: 'edit-rating-host/:id',
    component: EditRatingHostComponent
  },
  {
    path: 'show-ratings',
    component: ShowRatingsForHostComponent
  },
  {
    path: 'ratings-accommodation-by-guest',
    component: RatingsAccommodationByGuestComponent
  },
  {
    path: 'edit-rating-accommodation/:id',
    component: EditRatingAccommodationComponent
  },


];
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
