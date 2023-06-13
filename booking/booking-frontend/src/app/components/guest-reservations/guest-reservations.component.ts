import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import { MatTableDataSource } from "@angular/material/table";
import { Reservation } from 'src/app/model/reservation.model';
import { ReservationService } from 'src/app/service/reservation.service';

@Component({
  selector: 'app-guest-reservations',
  templateUrl: './guest-reservations.component.html',
  styleUrls: ['./guest-reservations.component.css']
})
export class GuestReservationsComponent implements OnInit {

  public dataSource = new MatTableDataSource<Reservation>();
  public displayedColumns = ['From', 'To', 'MaxNumberOfGuests', 'commands'];
  public accommodations: Reservation[] = [];
  public accommodation: Reservation | undefined = undefined;

  constructor() { }

  ngOnInit(): void {
  }

}
