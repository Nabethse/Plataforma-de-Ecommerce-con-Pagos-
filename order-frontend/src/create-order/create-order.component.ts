import { Component } from '@angular/core';
import { OrderService } from '../app/app.component';

@Component({
  selector: 'app-create-order',
  templateUrl: './create-order.component.html',
  styleUrls: ['./create-order.component.css']
})
export class CreateOrderComponent {
  orderData = { id: '', amount: 0, status: '' };
  response: any;

  constructor(private orderService: OrderService) {}

  createOrder() {
    this.orderService.createOrder(this.orderData).subscribe(
      (res) => {
        this.response = res;
      },
      (err) => {
        console.error(err);
      }
    );
  }
}