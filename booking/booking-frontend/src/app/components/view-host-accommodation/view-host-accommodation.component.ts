import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import { MatTableDataSource } from "@angular/material/table";
import Swal from 'sweetalert2';
import { User } from 'src/app/model/user.model';
import { Accommodation } from 'src/app/model/accommodation.model';
import { AccommodationService } from 'src/app/service/accommodation.service';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-view-host-accommodation',
  templateUrl: './view-host-accommodation.component.html',
  styleUrls: ['./view-host-accommodation.component.css']
})
export class ViewHostAccommodationComponent implements OnInit {

  public dataSource = new MatTableDataSource<Accommodation>();
  public displayedColumns = ['Name', 'MinNumberOfGuests', 'MaxNumberOfGuests', 'commands'];
  public accommodations: Accommodation[] = [];
  public accommodation: Accommodation | undefined = undefined;

  isLoggedIn: boolean = false;
  isHost: boolean = false;
  isGuest: boolean = false;
  user: User = new User();

  constructor(
    private router: Router,
    private accommodationService: AccommodationService,
    private userService : UserService
  ) { }

  ngOnInit(): void {
    let userRole = this.userService.getLoggedInUserRole()
    let userId = this.userService.getLoggedInUserId()
      this.accommodationService.getAccommodationByHostId(userId).subscribe((data: any) => {
      this.accommodations = data.accommodations;
      console.log(this.accommodations);
      this.dataSource.data = this.accommodations;
      console.log(this.dataSource.data);
    })

    if(userRole === "") {
      this.isLoggedIn = false;
    } else {
      this.isLoggedIn = true;
      if(userRole == "HOST") {
        this.isHost = true;
      }
      if(userRole == "GUEST") {
        this.isGuest = true;
      }
    }


  }

  makeAvailable(id: string){
    this.router.navigate(['createAvailability/' + id ]);
  }

}
