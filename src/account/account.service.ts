import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { authService, accountService } from 'src/db/account';

@Injectable()
export class AccountService {
  constructor(private jwtService: JwtService) {}

  async createUser(loginid: string, password: string) {
    const user = await accountService.createUser(loginid, password);
    return user;
  }

  async login(loginid: string, password: string) {
    const login = await authService.login(loginid, password);
    if (!login) {
      throw new Error('Invalid credentials');
    }

    const token = this.jwtService.sign({ userId: login.loginid });

    return { accessToken: token, loginid: login.loginid };
  }
}
