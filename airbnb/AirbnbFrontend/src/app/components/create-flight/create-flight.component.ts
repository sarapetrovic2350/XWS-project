import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {FlightService} from "../../service/flight.service";
import {Flight} from "../../model/flight.model";
import Swal from 'sweetalert2';

@Component({
  selector: 'app-create-flight',
  templateUrl: './create-flight.component.html',
  styleUrls: ['./create-flight.component.css']
})
export class CreateFlightComponent implements OnInit {

  constructor(
    private router: Router,
    private flightService: FlightService
  ) { }

  title = 'Create a Flight';
  flight = new Flight();
  submitted = false;
  date: string= "";
  time: string="";
  ngOnInit(): void {
  }

  onSubmit() {

    this.flight.availableSeats = this.flight.totalNumberOfSeats; 
    this.flight.dateTime = new Date(); 
    console.log(this.flight.dateTime); 
    //this.flight.dateTime = this.date.concat(this.time.toString());
    console.log(this.flight); 

    this.flightService.createFlight(this.flight).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/login']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully registered to Airbnb!',
          })

        },
        error: (e) => {
          this.submitted = false;
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Email already exists.',
          })

        }

      });
  }

}
