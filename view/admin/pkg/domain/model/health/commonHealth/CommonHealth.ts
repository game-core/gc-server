import type {CommonHealthType} from "~/pkg/domain/model/health/commonHealth/CommonHealthType";

export type CommonHealth = {
  "healthId": number;
  "name": string;
  "commonHealthType": CommonHealthType;
}
