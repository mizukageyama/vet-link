import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:frontend/features/pet/domain/breed/breed_model.dart';
import 'package:frontend/features/pet/domain/species/species_model.dart';

part 'pet_model.freezed.dart';
part 'pet_model.g.dart';

@Freezed()
class Pet with _$Pet {
  factory Pet({
    required int id,
    required String name,
    required String gender,
    required String birthdate,
    required List<Breed> breeds,
    required Species species,
  }) = _Pet;

  factory Pet.fromJson(Map<String, dynamic> json) => _$PetFromJson(json);
}
