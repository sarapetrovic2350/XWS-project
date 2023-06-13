import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import { Observable } from 'rxjs';
import {Flight} from "../model/flight.model";

@Injectable({
  providedIn: 'root'
})
export class FlightService {

  apiHost: string = 'http://localhost:8080/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json'});

  constructor(private http: HttpClient) { }

  createFlight(flight: Flight) {
    console.log(flight);
    return this.http.post<Flight>(this.apiHost + 'flights/createFlight', flight);
  }

  getAllFlights(): Observable<any[]> {
    return this.http.get<any[]>(this.apiHost + 'flights/getAllFlights');
  }

  deleteFlight(id: any): Observable<any> {
    return this.http.delete<any>(this.apiHost + 'flights/deleteFlight/' + id);
  }

  getFlightById(id: any): Observable<Flight>{
    return this.http.get<Flight>(this.apiHost + 'flights/' + id);
  }

  searchFlights(searchFlights: any): Observable<any[]>{
    return this.http.post<Flight[]>(this.apiHost + 'flights/searchFlights', searchFlights);
  }

  cancelFlight(id: any): Observable<any>{
    return this.http.put<any>(this.apiHost + 'flights/cancelFlight/' + id, {headers: this.headers});
  }

}
