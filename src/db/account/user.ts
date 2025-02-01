import prisma from './../db.service';
import { Account } from './account';
import { Password } from './password';

export class User {
  async checkUser(loginid: string): Promise<Account | null> {
    const user = await prisma.user.findUnique({ where: { loginid } });
    return user ? new Account(user.id, user.loginid, user.password) : null;
  }

  async createUser(loginid: string, password: string): Promise<Account> {
    const hashdPassword = await Password.hash(password);
    const user = await prisma.user.create({
      data: { loginid, password: hashdPassword },
    });
    return { id: user.id, loginid: user.loginid } as Omit<Account, 'password'>;
  }
}
