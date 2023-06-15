import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user.model';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-show-hosts',
  templateUrl: './show-hosts.component.html',
  styleUrls: ['./show-hosts.component.css']
})
export class ShowHostsComponent implements OnInit {
  
  public dataSource = new MatTableDataSource<User>();
  public displayedColumns = ['Name', 'Surname', 'Email', 'Address'];
  public users: User[] = [];
  public user: User | undefined = undefined;

  isLoggedIn: boolean = false;
  isHost: boolean = false;
  isGuest: boolean = false;
  //user: User = new User();

  constructor(
    private router: Router,
    //private accommodationService: AccommodationService,
    private userService : UserService
  ) { }

  ngOnInit(): void {
    let userRole = this.userService.getLoggedInUserRole()
    let userId = this.userService.getLoggedInUserId()
      this.userService.getAllHosts().subscribe((data: any) => {
      this.users = data.users;
      console.log(this.users);
      this.dataSource.data = this.users;
      console.log(this.dataSource.data);
    })


  }

}
