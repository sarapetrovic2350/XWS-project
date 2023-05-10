import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {FlightService} from "../../service/flight.service";
import {MatTableDataSource} from "@angular/material/table";
import {ShowFlight} from "../../model/show-flight.model";
import {Flight} from "../../model/flight.model";
import { User } from 'src/app/model/user.model';
import { UserService } from 'src/app/service/user.service';
import {AccommodationService} from "../../service/accommodation.service";
import {Accommodation} from "../../model/accommodation.model";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  //path: string = "../assets/images/plane.jpg";
  //alttext: string="image";

  constructor(private accommodationService: AccommodationService, private router: Router, private userService: UserService) {}
  startDate: Date = new Date()
  endDate: Date = new Date()
  country: string = ''
  city: string = ''
  numberOfGuests: number = 1
  public dataSource = new MatTableDataSource<Accommodation>();
  public displayedColumns = ['Name', 'MinNumberOfGuests', 'MaxNumberOfGuests', 'Address', 'Benefits'];
  public accommodations: Accommodation[] = [];
  public notFoundAccommodations: Accommodation[] = [];
  public accommodation: Accommodation = new Accommodation();
  isSearched: boolean = false;
  notFound: boolean = false;
  totalPrice: number = 0;
  public user: User = new User();
  role: string = "";

  ngOnInit(): void {
    //this.role = this.userService.getLoggedInUserRole();
  }
  searchAccommodations() {
    console.log(this.startDate)
    console.log(this.endDate)
    console.log(this.country)
    console.log(this.city)
    console.log(this.numberOfGuests)

    var newDate1 = new Date(this.startDate)
    console.log(newDate1)
    var newDate2 = new Date(newDate1.getFullYear(), newDate1.getMonth(), newDate1.getDate(), 2, 0, 0)
    console.log(newDate2)

    var newDate12 = new Date(this.endDate)
    console.log(newDate12)
    var newDate22 = new Date(newDate12.getFullYear(), newDate12.getMonth(), newDate12.getDate(), 2, 0, 0)
    console.log(newDate22)

    var searchAccommodations = {
      startDate: newDate1.toISOString(),
      endDate: newDate12.toISOString(),
      country: this.country,
      city: this.city,
      numberOfGuests: this.numberOfGuests
    }

    this.accommodationService.searchAccommodations(searchAccommodations).subscribe(
      {
        next: (res) => {
          console.log(res)
          this.isSearched = true;
          this.notFound = false;
          this.accommodations = res;
          for (let i = 0; i < this.accommodations.length; i++) {
            let startDtae = new Date(this.accommodations[i].startDate)
            this.accommodations[i].startDate = startDtae.toUTCString().replace('GMT', '')
            let endDate = new Date(this.accommodations[i].endDate)
            this.accommodations[i].endDate = endDate.toUTCString().replace('GMT', '')
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

  // public buyTicket(id: string) {
  //   if(this.user == null){
  //     this.router.navigate(['/login']);
  //   }else if (this.user.role == "REGISTERED_USER"){
  //     this.router.navigate(['createTicket/' + id ]);
  //   }
  // }


}
