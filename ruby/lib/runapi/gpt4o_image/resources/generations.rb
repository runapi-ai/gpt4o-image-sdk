# frozen_string_literal: true

module RunApi
  module Gpt4oImage
    module Resources
      class Generations
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/gpt4o_image/generations"

        RESPONSE_CLASS = Types::GenerationResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedGenerationResponse

        def initialize(http)
          @http = http
        end

        def run(**params)
          task = create(**params)
          poll_until_complete { get(task.id) }
        end

        def create(**params)
          params = compact_params(params)
          validate_params!(params)
          request(:post, ENDPOINT, body: params)
        end

        def get(id)
          request(:get, "#{ENDPOINT}/#{id}")
        end

        private

        def validate_params!(params)
          raise Core::ValidationError, "model is required" unless param(params, :model)
          raise Core::ValidationError, "size is required" unless param(params, :size)

          model = param(params, :model)
          unless Types::MODELS.include?(model)
            raise Core::ValidationError, "Invalid model: #{model}. Must be: #{Types::MODELS.join(", ")}"
          end

          validate_optional!(params, :size, Types::SIZES)
          validate_optional!(params, :n_variants, Types::N_VARIANTS)
          validate_optional!(params, :fallback_model, Types::FALLBACK_MODELS)
        end
      end
    end
  end
end
