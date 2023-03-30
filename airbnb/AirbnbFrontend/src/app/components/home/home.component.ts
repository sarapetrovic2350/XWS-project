import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  constructor() { }
  date: Date = new Date()
  departure: string = ''
  arrival: string = ''
  numberOfPassengers: number = 1

  ngOnInit(): void {
  }
  searchFlights() {
    console.log(this.date)
    console.log(this.departure)
    console.log(this.arrival)
    console.log(this.numberOfPassengers)
    

  }

}
