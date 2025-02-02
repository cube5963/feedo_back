import { Mutation, Resolver, Args, Query } from '@nestjs/graphql';
import { AccountService } from './account.service';
import { AuthResponse } from './dto/auth-response';
import { LoginInput } from './dto/login.input';
import { User } from './dto/user.model';
import { CreateUserInput } from './dto/createUser.dto';
import { UserSurvey } from './dto/query';

@Resolver()
export class AccountResolver {
  constructor(private readonly accountService: AccountService) {}

  /**
   *
   * @returns ユーザーがいるかどうかを確認するテストクエリのレスポンス
   */
  @Query(() => [UserSurvey])
  async getUserSurvey(@Args('loginid') loginid: string): Promise<UserSurvey[]> {
    return this.accountService.getUser(loginid);
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
