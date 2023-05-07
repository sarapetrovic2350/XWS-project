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
  departureDateDateForm: Date = new Date();
  arrivalDateDateForm: Date = new Date();
  departureTime: string = '';
  arrivalTime: string = '';
  ngOnInit(): void {
  }
  onSubmit() {

    let [departureHour, departureMinutes] = this.departureTime.split(':')
    this.departureDateDateForm.setHours( parseInt(departureHour)+2,  parseInt(departureMinutes), 0)
    console.log(this.departureDateDateForm)
    let [arrivalHour, arrivalMinutes] = this.arrivalTime.split(':')
    this.arrivalDateDateForm.setHours( parseInt(arrivalHour)+2,  parseInt(arrivalMinutes), 0)
    console.log(this.arrivalDateDateForm)

    this.flight.availableSeats = this.flight.totalNumberOfSeats;
    this.flight.departureDateTime = this.departureDateDateForm
    this.flight.arrivalDateTime = this.arrivalDateDateForm
    console.log(this.flight);

    this.flightService.createFlight(this.flight).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/showFlights']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully created new flight!',
          })

        },
        error: (e) => {
          this.submitted = false;
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
