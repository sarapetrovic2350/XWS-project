import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {FlightService} from "../../service/flight.service";
import {MatTableDataSource} from "@angular/material/table";
import {ShowFlight} from "../../model/show-flight.model";
import {Flight} from "../../model/flight.model";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  path: string = "../assets/images/plane.jpg";
  alttext: string="image";

  constructor(private flightService: FlightService, private router: Router) {}
  date: Date = new Date()
  departure: string = ''
  arrival: string = ''
  availableSeats: number = 1
  public dataSource = new MatTableDataSource<ShowFlight>();
  public displayedColumns = ['Departure', 'Arrival', 'DateTimeDeparture', 'DateTimeArrival', 'Price', 'TotalPrice'];
  public flights: ShowFlight[] = [];
  public flight: Flight | undefined = undefined;
  isSearched: boolean = false;
  totalPrice: number = 0;

  ngOnInit(): void {
  }
  searchFlights() {
    console.log(this.date)
    console.log(this.departure)
    console.log(this.arrival)
    console.log(this.availableSeats)

    var newDate1 = new Date(this.date)
    console.log(newDate1)
    var newDate2= new Date(newDate1.getFullYear(), newDate1.getMonth(), newDate1.getDate(), 2, 0, 0)
    console.log(newDate2)

    var searchFlights = { date:newDate1.toISOString(),
                          departure:this.departure,
                          arrival:this.arrival,
                          availableSeats:this.availableSeats
                           }

    this.flightService.searchFlights(searchFlights).subscribe((data: any)  =>
      {
          this.isSearched = true;
          this.flights = data;
          for (let i = 0; i < this.flights.length; i++) {
            let dateOfDeparture = new Date(this.flights[i].departureDateTime)
            this.flights[i].departureDateTime = dateOfDeparture.toUTCString().replace('GMT', '')
            let dateOfArrival = new Date(this.flights[i].arrivalDateTime)
            this.flights[i].arrivalDateTime = dateOfArrival.toUTCString().replace('GMT', '')
            this.flights[i].totalPrice = this.availableSeats*this.flights[i].price
          }
          this.dataSource.data = this.flights;
          console.log(data)
        })
      };

}
