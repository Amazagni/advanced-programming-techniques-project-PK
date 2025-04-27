import {bootstrapApplication} from '@angular/platform-browser';
import {AppComponent} from './app/app.component';
import {provideRouter} from '@angular/router';
import {HomeViewComponent} from './app/views/home-view/home-view.component';
import {InventoryViewComponent} from './app/views/inventory-view/inventory-view.component';
import {AddItemViewComponent} from './app/views/add-item-view/add-item-view.component';
import {provideHttpClient, withInterceptorsFromDi} from '@angular/common/http';

bootstrapApplication(AppComponent, {
  providers: [
    provideRouter([
      {path: '', component: HomeViewComponent},
      {path: 'inventory', component: InventoryViewComponent},
      {path: 'add-item', component: AddItemViewComponent},
    ]),
    provideHttpClient(withInterceptorsFromDi())
  ]
}).catch(err => console.error(err));
