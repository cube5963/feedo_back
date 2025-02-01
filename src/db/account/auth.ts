import { User } from './user';
import { Password } from './password';
import { Account } from './account';

export class Auth {
  private account: User;

  constructor() {
    this.account = new User();
  }

  async login(loginid: string, password: string): Promise<Account | null> {
    const user = await this.account.checkUser(loginid);
    if (!user) return null;

    const isValid = await Password.compare(password, user.getPassword());
    return isValid ? user : null;
  }
}
