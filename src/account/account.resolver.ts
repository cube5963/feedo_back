import { Mutation, Resolver, Args, Query } from '@nestjs/graphql';
import { AccountService } from './account.service';
import { AuthResponse } from './dto/auth-response';
import { LoginInput } from './dto/login.input';
import { User } from './dto/user.model';
import { CreateUserInput } from './dto/createUser.dto';

@Resolver()
export class AccountResolver {
  constructor(private readonly accountService: AccountService) {}

  /**
   *
   * @returns ユーザーがいるかどうかを確認するテストクエリのレスポンス
   */
  @Query(() => String)
  async testQuery() {
    return 'Test query response';
  }

  @Mutation(() => User)
  async createUser(
    @Args('createuserinput') CreateUserInput: CreateUserInput,
  ): Promise<User> {
    return this.accountService.createUser(
      CreateUserInput.loginid,
      CreateUserInput.password,
    );
  }

  @Mutation(() => AuthResponse)
  async login(
    @Args('loginInput') loginInput: LoginInput,
  ): Promise<AuthResponse> {
    return this.accountService.login(loginInput.loginid, loginInput.password);
  }
}
