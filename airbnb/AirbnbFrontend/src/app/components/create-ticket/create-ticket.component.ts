import { Component, OnInit } from '@angular/core';
import Swal from 'sweetalert2';
import {Router} from "@angular/router";
import { Ticket } from 'src/app/model/ticket.model';
import { Flight } from 'src/app/model/flight.model';
import { User } from 'src/app/model/user.model';
import { TicketService } from 'src/app/service/ticket.service';
import { UserService } from 'src/app/service/user.service';
import { FlightService } from 'src/app/service/flight.service';

@Component({
  selector: 'app-create-ticket',
  templateUrl: './create-ticket.component.html',
  styleUrls: ['./create-ticket.component.css']
})
export class CreateTicketComponent implements OnInit {

  public user: User = new User(); 
  public ticket: Ticket = new Ticket(); 
  public flight: Flight = new Flight(); 
  submitted = false;

  public title = "Buy Tickets"

  public flightId = "64240576e3f0e3cfa51d1257"; 

  constructor(
    private flightService: FlightService, 
    private router: Router,
    private userService: UserService, 
    private ticketService: TicketService
    ) { }

  ngOnInit(): void {
    this.user = this.userService.getCurrentUser(); 
    
    this.flightService.getFlightById(this.flightId).subscribe(res => { 
      this.flight = res; 
      console.log(this.flight); 
    })

  }

  public createTicket(){

    this.ticket.idFlight = this.flight.id; 
    this.ticket.idUser = this.user.id; 

    console.log(this.ticket); 

    this.ticketService.createTicket(this.ticket).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Get packing!',
          })

        },
        error: (e) => {
          this.submitted = false;
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Sorry, not enough free seats.',
          })

        }

      });

  }

}
