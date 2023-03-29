import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import { MatTableDataSource } from "@angular/material/table";
import { User } from 'src/app/model/user.model';
import { Ticket } from 'src/app/model/ticket.model';
import { Flight } from 'src/app/model/flight.model';
import { ShowTicket } from 'src/app/model/show-ticket.model';
import { TicketService } from 'src/app/service/ticket.service'; 
import { UserService } from 'src/app/service/user.service';
import { FlightService } from 'src/app/service/flight.service'; 

@Component({
  selector: 'app-show-users-tickets',
  templateUrl: './show-users-tickets.component.html',
  styleUrls: ['./show-users-tickets.component.css']
})
export class ShowUsersTicketsComponent implements OnInit {

  public dataSource = new MatTableDataSource<ShowTicket>();
  public displayedColumns = ['DateOfPurchase', 'Departure', 'Arrival', 'DateOfDeparture','NumberOfTickets', 'TotalPrice' ];
  public tickets: ShowTicket[] = [];
  public ticket: ShowTicket | undefined = undefined;

  public user: User = new User(); 

  constructor(
    private flightService: FlightService, 
    private router: Router,
    private userService: UserService, 
    private ticketService: TicketService
  ) { }

  ngOnInit(): void {

    this.user = this.userService.getCurrentUser(); 
    
    this.ticketService.getTicketsByUserId(this.user.id).subscribe((data: any) => { 
      this.tickets = data; 
      this.dataSource.data = this.tickets;
      console.log(this.tickets); 
    })
  }

}
