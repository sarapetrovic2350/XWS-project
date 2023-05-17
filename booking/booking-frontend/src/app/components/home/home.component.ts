import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import Swal from 'sweetalert2';
import {FlightService} from "../../service/flight.service";
import {MatTableDataSource} from "@angular/material/table";
import {ShowFlight} from "../../model/show-flight.model";
import {Flight} from "../../model/flight.model";
import { User } from 'src/app/model/user.model';
import { UserService } from 'src/app/service/user.service';
import {AccommodationService} from "../../service/accommodation.service";
import {Accommodation} from "../../model/accommodation.model";
import { Availability } from 'src/app/model/availability.model';

import { Reservation } from 'src/app/model/reservation.model';
import { ReservationService } from 'src/app/service/reservation.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  //path: string = "../assets/images/plane.jpg";
  //alttext: string="image";

  constructor(private reservationService: ReservationService, private accommodationService: AccommodationService, private router: Router, private userService: UserService) {}
  startDate: Date = new Date()
  endDate: Date = new Date()
  country: string = ''
  city: string = ''
  numberOfGuests: number = 1
  public dataSource = new MatTableDataSource<Accommodation>();

  public displayedColumns = ['Name', 'MinNumberOfGuests', 'MaxNumberOfGuests', 'Address', 'Benefits', 'Status','Price', 'commands'];

  public accommodations: Accommodation[] = [];
  public notFoundAccommodations: Accommodation[] = [];
  public accommodation: Accommodation = new Accommodation();
  isSearched: boolean = false;
  notFound: boolean = false;
  public totalPrice: number = 0;
  public user: User = new User();
  role: string = "";
  public price : number = 0;
  public priceSelection: string = '';
  public availabilities: Availability[] = [];

  ngOnInit(): void {
    //this.role = this.userService.getLoggedInUserRole();
  }
  searchAccommodations() {
    console.log(this.startDate)
    console.log(this.endDate)
    console.log(this.country)
    console.log(this.city)
    console.log(this.numberOfGuests)


   var searchParams
    var searchAccommodations = {
        searchParams : {
          country: this.country,
          city: this.city,
          numberOfGuests: this.numberOfGuests,
          startDate: this.startDate,
          endDate: this.endDate,
      }
    }
    console.log(searchAccommodations)


    this.accommodationService.searchAccommodations(searchAccommodations).subscribe(
      {
        next: (res) => {
          console.log(res)
          this.isSearched = true;
          this.notFound = false;
          this.accommodations = res.accommodations;
          for (let i = 0; i < this.accommodations.length; i++) {
            //let startDate = new Date(this.accommodations[i].startDate)
            //this.accommodations[i].startDate = startDate.toUTCString().replace('GMT', '')
            //let endDate = new Date(this.accommodations[i].endDate)
            //this.accommodations[i].endDate = endDate.toUTCString().replace('GMT', '')
            console.log(this.accommodations[i].availabilities)
            this.availabilities = this.accommodations[i].availabilities
            var startDate1 = this.startDate;
            var endDate1 = this.endDate;
            for (let i = 0; i < this.availabilities.length; i++){
              this.price  = this.availabilities[i].price
              this.priceSelection = this.availabilities[i].priceSelection.toString() 
              console.log(this.price)
              console.log(this.priceSelection)
              if(this.priceSelection == "PER_PERSON"){
               // console.log(this.endDate.getTime())
                //var sub = endDate1.getDate() - startDate1.getDate()
                //console.log(this.endDate)
                //console.log(this.startDate)
                //console.log(this.startDate.getDay())
                //this.totalPrice = sub*this.price*this.numberOfGuests
                //console.log(sub)
                //console.log(this.totalPrice)
              }else{
                //var sub = this.endDate.getTime() - this.startDate.getTime()
                //this.totalPrice = sub*this.price
              }

              //this.priceSelection = this.accommodation.availabilities[i].priceSelection
            }
          }
          this.dataSource.data = this.accommodations;
          console.log(this.accommodations)

        },

        error: (e) => {
          this.notFound = true;
          this.isSearched = true;
          this.dataSource.data = this.notFoundAccommodations;
          console.log(e);
        }
      });
  }

  clearSearch() {
    this.country= ''
    this.city = ''
    this.numberOfGuests= 1
    this.startDate = new Date()
    this.endDate = new Date()
    this.isSearched = false;
    this.notFound = false;
  }

  reserve(id: string){
    let userId = this.userService.getLoggedInUserId();

    var NewReservation = {
      numberOfGuests: this.numberOfGuests,
      startDate: this.startDate,
      endDate: this.endDate,
      userId: userId,
      accommodationId: id
    }

    console.log(this.startDate);
    console.log(this.endDate);

    this.reservationService.createReservation(NewReservation).subscribe(
      {
        next: (res) => {
          //this.router.navigate(['/show-host-accommodations']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully created new accommodation!',
          })

        },
        error: (e) => {

          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'There are already reservations in that period.',
          })

        }

      });

  }

  // public buyTicket(id: string) {
  //   if(this.user == null){
  //     this.router.navigate(['/login']);
  //   }else if (this.user.role == "REGISTERED_USER"){
  //     this.router.navigate(['createTicket/' + id ]);
  //   }
  // }


}
