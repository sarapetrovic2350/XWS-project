import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import Swal from 'sweetalert2';
import {MatTableDataSource} from "@angular/material/table";
import { Reservation } from 'src/app/model/reservation.model';
import { ReservationService } from 'src/app/service/reservation.service';
import { User } from 'src/app/model/user.model';
import { UserService } from 'src/app/service/user.service';
import {AccommodationService} from "../../service/accommodation.service";
import {Accommodation} from "../../model/accommodation.model";

@Component({
  selector: 'app-view-reservations',
  templateUrl: './view-reservations.component.html',
  styleUrls: ['./view-reservations.component.css']
})
export class ViewReservationsComponent implements OnInit {

  constructor(
    private reservationService: ReservationService, 
    private router: Router, 
    private userService: UserService, 
    private accommodationService: AccommodationService) 
  {}
  startDate: Date = new Date()
  endDate: Date = new Date()
  public dataSource = new MatTableDataSource<Reservation>();

  public displayedColumns = ['From', 'To', 'City', 'Country', 'AccommodationsName', 'Status'];

  public reservations: Reservation[] = [];
  public reservation: Reservation = new Reservation();

  

  ngOnInit(): void {
    let userId = this.userService.getLoggedInUserId();

    this.reservationService.getAllReservationsByGuestId(userId).subscribe(
      {
        next: (data) => {
          
          this.reservations = data;
          for (let i = 0; i < this.reservations.length; i++) {
            let startDate = new Date(this.reservations[i].startDate)
            this.reservations[i].startDate = startDate.toUTCString().replace('GMT', '')
            let endDate = new Date(this.reservations[i].endDate)
            this.reservations[i].endDate = endDate.toUTCString().replace('GMT', '')

            this.accommodationService.getAccommodationById(this.reservations[i].accommodationId).subscribe({
              next: (res: Accommodation) => {
                this.reservations[i].name = res.name; 
                this.reservations[i].city = res.address.city; 
                this.reservations[i].country = res.address.country; 
              }
            }); 
        
          }
          this.dataSource.data = this.reservations;

          console.log(this.reservations)

        },

        error: (e) => {
          // this.notFound = true;
          // this.isSearched = true;
          // this.dataSource.data = this.notFoundAccommodations;
          // console.log(e);
        }
      });
  }

}
