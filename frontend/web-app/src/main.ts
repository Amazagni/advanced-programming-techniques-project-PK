import {bootstrapApplication} from '@angular/platform-browser';
import {AppComponent} from './app/app.component';
import {provideRouter} from '@angular/router';
import {HomeViewComponent} from './app/views/home-view/home-view.component';
import {InventoryViewComponent} from './app/views/inventory-view/inventory-view.component';
import {AddItemViewComponent} from './app/views/add-item-view/add-item-view.component';
import {provideHttpClient, withInterceptorsFromDi} from '@angular/common/http';
import {AdminGuard, AuthGuard} from './app/guards/auth.guard';
import {LoginViewComponent} from './app/views/login-view/login-view.component';

bootstrapApplication(AppComponent, {
  providers: [
    provideRouter([
      { path: '', component: HomeViewComponent },
      { path: 'inventory', component: InventoryViewComponent, canActivate: [AuthGuard] },
      { path: 'add-item', component: AddItemViewComponent, canActivate: [AdminGuard] },
      { path: 'login', component: LoginViewComponent },
    ]),
    provideHttpClient(withInterceptorsFromDi())
  ]
}).catch(err => console.error(err));
