import { Mutation, Resolver, Args } from '@nestjs/graphql';
import { SurveyService } from './survey.service';
import { CreateSurveyInput } from './dto/create.dto';

@Resolver()
export class SurveyResolver {
  constructor(private readonly surveyService: SurveyService) {}

  @Mutation(() => Boolean)
  async createSurvey(
    @Args('createsurveyinput') CreateSurveyInput: CreateSurveyInput,
  ): Promise<boolean> {
    await this.surveyService.Create(CreateSurveyInput);
    return true;
  }
}
