import { Field, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class UserSurvey {
  @Field()
  id: number;

  @Field()
  title: string;
}
