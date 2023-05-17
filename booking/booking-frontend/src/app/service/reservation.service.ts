import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import { Observable } from 'rxjs';
import { Reservation } from '../model/reservation.model';

@Injectable({
  providedIn: 'root'
})
export class ReservationService {

  apiHost: string = 'http://localhost:8000/reservation';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Access-Control-Allow-Origin' : '*',
  'Access-Control-Allow-Methods' : 'GET,HEAD,OPTIONS,POST,PUT', 'Access-Control-Allow-Headers' : 'Origin, X-Requested-With, Content-Type, Accept, x-client-key, x-client-token, x-client-secret, Authorization' }  );
  headers2: HttpHeaders = new HttpHeaders({
  'Content-Type': 'application/json',
  });
  apiHost1: string = 'http://localhost:8000/';
  constructor(private http: HttpClient) { }

  createReservation(reservation: any) {
    console.log(reservation);
    return this.http.post<Reservation>(this.apiHost , reservation, {headers: this.headers2});
  }

  getActiveReservationsByGuestId(id: any){
    return this.http.get<any[]>(this.apiHost + 'getActiveReservationsByGuestId/' + id, {headers: this.headers2});
  }

  getAllReservationsByGuestId(id: any){
    return this.http.get<any>(this.apiHost1 + 'getReservationsByUserId/' + id, {headers: this.headers2});
  }

  cancelReservation(id: string){
    return this.http.post<Reservation>(this.apiHost + '/cancelReservationByGuest/' + id, {headers: this.headers2});
  }

  deleteReservation(id: string){
    return this.http.post<Reservation>(this.apiHost + '/deletePendingReservation/' + id, {headers: this.headers2});
  }
}
