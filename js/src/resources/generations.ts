import type { HttpClient, PollingOptions, RequestOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import type { GenerationParams, GenerationResponse, TaskCreateResponse } from '../types';

const ENDPOINT = '/api/v1/gpt4o_image/generations';

export class Generations {
  constructor(private readonly http: HttpClient) {}

  async run(params: GenerationParams, options?: RequestOptions & PollingOptions): Promise<GenerationResponse> {
    const { id } = await this.create(params, options);
    return pollUntilComplete<GenerationResponse>(() => this.get(id, options), {
      maxWaitMs: options?.maxWaitMs,
      pollIntervalMs: options?.pollIntervalMs,
    });
  }

  async create(params: GenerationParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    return this.http.request<TaskCreateResponse>('POST', ENDPOINT, {
      body: compactParams(params),
      ...options,
    });
  }

  async get(id: string, options?: RequestOptions): Promise<GenerationResponse> {
    return this.http.request<GenerationResponse>('GET', `${ENDPOINT}/${id}`, {
      ...options,
    });
  }
}
