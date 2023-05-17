import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AngularMaterialModule } from './angular-material/angular-material.module';
import { UserRegistrationComponent } from './components/user-registration/user-registration.component';
import { HeaderComponent } from './components/header/header.component';
import {HttpClientModule} from "@angular/common/http";
import { LoginComponent } from './components/login/login.component';
import { ShowFlightsComponent } from './components/show-flights/show-flights.component';
import { CreateFlightComponent } from './components/create-flight/create-flight.component';
import {ReactiveFormsModule} from "@angular/forms";
import { HomeComponent } from './components/home/home.component';
import { CreateAccommodationComponent } from './components/create-accommodation/create-accommodation.component';
import { ViewHostAccommodationComponent } from './components/view-host-accommodation/view-host-accommodation.component';
import { CreateAvailabilityComponent } from './components/create-availability/create-availability.component';
import { UpdateUserComponent } from './components/update-user/update-user.component';
import { GuestReservationsComponent } from './components/guest-reservations/guest-reservations.component';
import { ViewReservationsComponent } from './components/view-reservations/view-reservations.component';
import { ViewPendingReservationsComponent } from './components/view-pending-reservations/view-pending-reservations.component';



@NgModule({
  declarations: [
    AppComponent,
    UserRegistrationComponent,
    HeaderComponent,
    LoginComponent,
    ShowFlightsComponent,
    CreateFlightComponent,
    HomeComponent,
    CreateAccommodationComponent,
    ViewHostAccommodationComponent,
    CreateAvailabilityComponent,
    UpdateUserComponent,
    GuestReservationsComponent,
    ViewReservationsComponent,
    ViewPendingReservationsComponent
  ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        AngularMaterialModule,
        HttpClientModule,
        ReactiveFormsModule
    ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
