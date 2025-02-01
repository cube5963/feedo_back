export class Account {
  constructor(
    public id: number,
    public loginid: string,
    private password: string,
  ) {}

  getPassword() {
    return this.password;
  }
}
