import { Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import * as bcrypt from 'bcrypt';

@Injectable()
export class AccountService {
  constructor(private jwtService: JwtService) {}

  async login(loginid: string, password: string) {
    const tpass = await bcrypt.hash('test', 10);
    const user = { loginid: 'testuser', passwordHash: tpass };

    if (!user) {
      throw new Error('User not found');
    }

    const isPasswordMatching = await bcrypt.compare(
      password,
      user.passwordHash,
    );

    if (!isPasswordMatching) {
      throw new Error('Invalid credentials');
    }

    const token = this.jwtService.sign({ userId: user.loginid });

    return { accessToken: token, loginid: user.loginid };
  }
}
