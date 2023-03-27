import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import { MatTableDataSource } from "@angular/material/table";
import {FlightService} from "../../service/flight.service";
import {Flight} from "../../model/flight.model";

@Component({
  selector: 'app-show-flights',
  templateUrl: './show-flights.component.html',
  styleUrls: ['./show-flights.component.css']
})
export class ShowFlightsComponent implements OnInit {

  public dataSource = new MatTableDataSource<Flight>();
  public displayedColumns = ['DateTime', 'Departure', 'Arrival', 'Price', 'TotalNumberOfSeats','AvailableSeats'];
  public flights: Flight[] = [];
  public flight: Flight | undefined = undefined;

  constructor(private flightService: FlightService, private router: Router) { }

  ngOnInit(): void {

    this.flightService.getAllFlights().subscribe((data: any) => {  
      console.log(this.flights);     
      this.flights = data;
      this.dataSource.data = this.flights;
    })
    

  }

}
