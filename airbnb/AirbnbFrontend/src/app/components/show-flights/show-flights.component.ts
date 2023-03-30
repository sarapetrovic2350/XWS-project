import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import { MatTableDataSource } from "@angular/material/table";
import {FlightService} from "../../service/flight.service";
import {Flight} from "../../model/flight.model";
import Swal from 'sweetalert2';

@Component({
  selector: 'app-show-flights',
  templateUrl: './show-flights.component.html',
  styleUrls: ['./show-flights.component.css']
})
export class ShowFlightsComponent implements OnInit {

  public dataSource = new MatTableDataSource<Flight>();
  public displayedColumns = ['DateTime', 'Departure', 'Arrival', 'Price', 'TotalNumberOfSeats','AvailableSeats', 'commands'];
  public flights: Flight[] = [];
  public flight: Flight | undefined = undefined;

  public date: Date = new Date()
  public departure: string = ''
  public arrival: string = ''
  public availableSeats: number = 1

  constructor(private flightService: FlightService, private router: Router) { }

  ngOnInit(): void {

    this.flightService.getAllFlights().subscribe((data: any) => {  
      console.log(this.flights);     
      this.flights = data;
      this.dataSource.data = this.flights;
    })
    

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
          this.flights = res;
          this.dataSource.data = this.flights;
        })   
      };

  public deleteFlight(id: string) {
    console.log(id)
    this.flightService.deleteFlight(id).subscribe(res =>
      {
        this.flightService.getAllFlights().subscribe(res => {     
          this.flights = res;
          this.dataSource.data = this.flights;
        })
        //window.location.reload(); 
        

      });
  }

}
