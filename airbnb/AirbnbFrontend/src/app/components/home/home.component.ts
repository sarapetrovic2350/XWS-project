import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {FlightService} from "../../service/flight.service";

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
  // model koji ce se koristiti za prikaz
  //public flights: ShowFlight[] = [];

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

    this.flightService.searchFlights(searchFlights).subscribe(res =>
      {   
          // ovo otkomentarisati
          //this.flights = res;
          //this.dataSource.data = this.flights;
          console.log(res)
        })   
      };

}
