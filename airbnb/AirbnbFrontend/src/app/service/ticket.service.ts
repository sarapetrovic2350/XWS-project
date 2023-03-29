import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import { Ticket } from '../model/ticket.model';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TicketService {

  apiHost: string = 'http://localhost:8080/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json'});

  constructor(private http: HttpClient) { }

  createTicket(ticket: Ticket) {
    console.log(ticket); 
    return this.http.post<Ticket>(this.apiHost + 'tickets/createTicket', ticket);
  }

}
