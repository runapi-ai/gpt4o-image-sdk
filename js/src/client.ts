import { createHttpClient, type ClientOptions } from '@runapi.ai/core';
import { Generations } from './resources/generations';

/**
 * Gpt4o Image generation API client.
 */
export class Gpt4oImageClient {
  public readonly generations: Generations;

  constructor(options: ClientOptions = {}) {
    const http = createHttpClient(options);
    this.generations = new Generations(http);
  }
}
