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
  selector: 'app-view-pending-reservations',
  templateUrl: './view-pending-reservations.component.html',
  styleUrls: ['./view-pending-reservations.component.css']
})
export class ViewPendingReservationsComponent implements OnInit {

  constructor(
    private reservationService: ReservationService, 
    private router: Router, 
    private userService: UserService, 
    private accommodationService: AccommodationService) 
  {}

  startDate: Date = new Date()
  endDate: Date = new Date()
  public dataSource = new MatTableDataSource<Reservation>();

  public displayedColumns = ['From', 'To', 'City', 'Country', 'AccommodationsName', 'Status', 'commands', 'commands1'];
  public reservations: Reservation[] = [];
  public reservation: Reservation = new Reservation();

  public accommodation: Accommodation = new Accommodation(); 

  
  ngOnInit(): void {
    let userId = this.userService.getLoggedInUserId();

    this.reservationService.getPendingReservationsByHostId(userId).subscribe(
      {
        next: (res) => {
          
          this.reservations = res.reservations;
          
          for (let i = 0; i < this.reservations.length; i++) {
            console.log(this.reservations[i].accommodationId); 
            // let startDate = new Date(this.reservations[i].startDate)
            // this.reservations[i].startDate = startDate.toUTCString().replace('GMT', '')
            // let endDate = new Date(this.reservations[i].endDate)
            // this.reservations[i].endDate = endDate.toUTCString().replace('GMT', '')

            console.log(this.reservations[i].accommodationId); 
            console.log(this.reservations[i].reservationStatus); 
            this.accommodationService.getAccommodationById(this.reservations[i].accommodationId).subscribe({
              next: (res) => {
                this.accommodation = res.accommodation;  

                this.reservations[i].name = this.accommodation.name; 
                this.reservations[i].city = this.accommodation.address.city; 
                this.reservations[i].country = this.accommodation.address.country; 
                

              }
            }); 
        
          }
          this.dataSource.data = res.reservations;

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

  accept(id: string){
    this.reservationService.AcceptPendingReservationByHost(id).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/view-reservations']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully accepted reservation!',
          })
          window.location.reload();
        },

        error: (e) => {
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong.',
          })
        }
      });
  }

  reject(id: string){
    this.reservationService.RejectPendingReservationByHost(id).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/view-reservations']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully rejected reservation!',
          })
          window.location.reload();
        },

        error: (e) => {
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong.',
          })
        }
      });
  }

}
