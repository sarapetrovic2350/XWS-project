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

  public displayedColumns = ['Name', 'MinNumberOfGuests', 'MaxNumberOfGuests', 'Address', 'Benefits', 'Status', 'Price', 'commands'];

  public accommodations: Accommodation[] = [];
  public notFoundAccommodations: Accommodation[] = [];
  public accommodation: Accommodation = new Accommodation();
  isSearched: boolean = false;
  notFound: boolean = false;
  public totalPrice: number = 0;
  public user: User = new User();
  role: string = "";
  public price: number = 0;
  public priceSelection: string = '';
  public availabilities: Availability[] = [];

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
    //this.userService.getUserByEmail(userEmail).subscribe(res => {
      //this.user = res;
      //console.log(this.user)
    //})

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


    var searchParams
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


    this.accommodationService.searchAccommodations(searchAccommodations).subscribe(
      {
        next: (res) => {
          console.log(res)
          this.isSearched = true;
          this.notFound = false;
          this.accommodations = res.accommodations;
          for (let i = 0; i < this.accommodations.length; i++) {
            this.availabilities = this.accommodations[i].availabilities
            var startDate1 = this.startDate;
            var endDate1 = this.endDate;
            for (let i = 0; i < this.availabilities.length; i++){
              this.price  = this.availabilities[i].price
              this.priceSelection = this.availabilities[i].priceSelection.toString()

            }
          }

          this.filterAccommodations();
          this.dataSource.data = this.filteredAccommodations;
          

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
  }

  filterAccommodations() {
    this.filteredAccommodations = this.accommodations.filter((acc: any) => {
      const price = acc.availabilities[0].price;
      const selectedBenefits = this.selectedBenefits;

      // Proveri da li accommodation sadrži sve selektovane benefite
      const hasSelectedBenefits = selectedBenefits.every((benefit: string) =>
        acc.benefits.some((accBenefit: any) => accBenefit == benefit)
      );

      return (
        (this.minPrice === undefined || price >= this.minPrice) &&
        (this.maxPrice === undefined || price <= this.maxPrice) &&
        (selectedBenefits.length === 0 || hasSelectedBenefits)
      );
    });
  }

  toggleBenefitSelection(benefit: string) {
    const index = this.selectedBenefits.indexOf(benefit);
    if (index > -1) {
      // Ako je benefit već selektovan, ukloni ga iz niza selectedBenefits
      this.selectedBenefits.splice(index, 1);
    } else {
      // Ako benefit nije selektovan, dodaj ga u niz selectedBenefits
      this.selectedBenefits.push(benefit);
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

    console.log(this.startDate);
    console.log(this.endDate);

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

