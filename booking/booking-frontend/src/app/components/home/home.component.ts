import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import Swal from 'sweetalert2';
import {MatTableDataSource} from "@angular/material/table";
import { User } from 'src/app/model/user.model';
import { UserService } from 'src/app/service/user.service';
import {AccommodationService} from "../../service/accommodation.service";
import {Accommodation} from "../../model/accommodation.model";
import { Availability } from 'src/app/model/availability.model';
import { ReservationService } from 'src/app/service/reservation.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  //path: string = "../assets/images/plane.jpg";
  //alttext: string="image";

  constructor(private reservationService: ReservationService, private accommodationService: AccommodationService, private router: Router, private userService: UserService) {
  }

  startDate: Date = new Date()
  endDate: Date = new Date()
  country: string = ''
  city: string = ''
  numberOfGuests: number = 1
  public dataSource = new MatTableDataSource<Accommodation>();

  public displayedColumns = ['Name', 'MinNumberOfGuests', 'MaxNumberOfGuests', 'Address', 'Benefits', 'Status', 'Price', 'Total Price', 'commands'];

  public accommodations: Accommodation[] = [];
  public notFoundAccommodations: Accommodation[] = [];
  public accommodation: Accommodation = new Accommodation();
  isSearched: boolean = false;
  notFound: boolean = false;
  public user: User = new User();
  role: string = "";
  public price: number = 0;
  public priceSelection: string = '';
  public availabilities: any[] = [];
  isLoggedIn: boolean = false;
  isHost: boolean = false;
  isGuest: boolean = false;

  public filteredAccommodations: Accommodation[] = [];
  public minPrice: number = 0;
  public maxPrice: number = Infinity;
  public selectedBenefits: string[] = [];
  public benefits: string[] = ['Wifi', 'Free Parking', 'Private Bathroom', 'Shared Bathroom', 'Kitchen',
  'Air Conditioner'];

  ngOnInit(): void {
    //this.role = this.userService.getLoggedInUserRole();
    let userRole = this.userService.getLoggedInUserRole()
    let userEmail = this.userService.getLoggedInUserEmail()
    console.log(userEmail)

    if(userRole === "") {
      this.isLoggedIn = false;
    } else {
      this.isLoggedIn = true;
      if(userRole == "HOST") {
        this.isHost = true;
      }
      if(userRole == "GUEST") {
        this.isGuest = true;
      }
    }
  }

  searchAccommodations() {
    console.log(this.startDate)
    console.log(this.endDate)
    console.log(this.country)
    console.log(this.city)
    console.log(this.numberOfGuests)

    var searchAccommodations = {
      searchParams: {
        country: this.country,
        city: this.city,
        numberOfGuests: this.numberOfGuests,
        startDate: this.startDate,
        endDate: this.endDate,
      }
    }
    console.log(searchAccommodations)

    let startDateString = this.startDate.toString()
    let endDateString = this.endDate.toString()
    let startDate = new Date(Date.parse(startDateString))
    let endDate = new Date(Date.parse(endDateString))

    this.accommodationService.searchAccommodations(searchAccommodations).subscribe(
      {
        next: (res) => {
          this.isSearched = true;
          this.notFound = false;
          this.accommodations = res.accommodations;
          for (let i = 0; i < this.accommodations.length; i++) {
            this.availabilities.push(this.accommodations[i].availabilities)
          }
          console.log(this.availabilities)
          for (let i = 0; i < this.accommodations.length; i++) {
            for (let i = 0; i < this.availabilities.length; i++) {
              this.accommodations[i].price = this.availabilities[i][0].price
              this.accommodations[i].priceSelection = this.availabilities[i][0].priceSelection
              let days = Math.floor((endDate.getTime() - startDate.getTime()) / (1000 * 3600 * 24));
              console.log(days)
              if (this.accommodations[i].priceSelection == "PER_PERSON") {

                this.accommodations[i].totalPrice = days * this.accommodations[i].price * this.numberOfGuests
              } else {
                this.accommodations[i].totalPrice = days * this.accommodations[i].price
              }
            }

          }

          this.dataSource.data = this.accommodations;

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
    this.country = ''
    this.city = ''
    this.numberOfGuests = 1
    this.startDate = new Date()
    this.endDate = new Date()
    this.isSearched = false;
    this.notFound = false;
    this.isSearched = false;
    this.availabilities = [];
  }

  filterAccommodations() {
    if(this.isSearched){
      this.filteredAccommodations = this.accommodations.filter((acc: any) => {
        const price = acc.availabilities[0].price;
        const selectedBenefits = this.selectedBenefits;
        const isSuperHost = acc.isSuperHost
        // Proveri da li accommodation sadrži sve selektovane benefite
        const hasSelectedBenefits = selectedBenefits.every((benefit: string) =>
          acc.benefits.some((accBenefit: any) => accBenefit == benefit)
        );

        return (
          (this.minPrice === undefined || price >= this.minPrice) &&
          (this.maxPrice === undefined || price <= this.maxPrice) &&
          (selectedBenefits.length === 0 || hasSelectedBenefits) &&
          isSuperHost
        );
      });
      this.dataSource.data = this.filteredAccommodations;
    }
  }

  reserve(id: string) {
    let userId = this.userService.getLoggedInUserId();

    var NewReservation = {
      numberOfGuests: this.numberOfGuests,
      startDate: this.startDate,
      endDate: this.endDate,
      userId: userId,
      accommodationId: id
    }

    this.reservationService.createReservation(NewReservation).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/view-reservations']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully made a reservation!',
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
}

