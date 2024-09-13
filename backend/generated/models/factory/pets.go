// Code generated by BobGen mysql v0.28.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"
	"testing"

	models "backend/generated/models"
	"github.com/aarondl/opt/omit"
	"github.com/jaswdr/faker/v2"
	"github.com/stephenafamo/bob"
)

type PetMod interface {
	Apply(*PetTemplate)
}

type PetModFunc func(*PetTemplate)

func (f PetModFunc) Apply(n *PetTemplate) {
	f(n)
}

type PetModSlice []PetMod

func (mods PetModSlice) Apply(n *PetTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// PetTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type PetTemplate struct {
	ID        func() uint32
	Name      func() string
	Gender    func() uint8
	Birthdate func() string
	BreedID   func() uint32
	OwnerID   func() uint32

	r petR
	f *Factory
}

type petR struct {
	Breed     *petRBreedR
	OwnerUser *petROwnerUserR
}

type petRBreedR struct {
	o *BreedTemplate
}
type petROwnerUserR struct {
	o *UserTemplate
}

// Apply mods to the PetTemplate
func (o *PetTemplate) Apply(mods ...PetMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.Pet
// this does nothing with the relationship templates
func (o PetTemplate) toModel() *models.Pet {
	m := &models.Pet{}

	if o.ID != nil {
		m.ID = o.ID()
	}
	if o.Name != nil {
		m.Name = o.Name()
	}
	if o.Gender != nil {
		m.Gender = o.Gender()
	}
	if o.Birthdate != nil {
		m.Birthdate = o.Birthdate()
	}
	if o.BreedID != nil {
		m.BreedID = o.BreedID()
	}
	if o.OwnerID != nil {
		m.OwnerID = o.OwnerID()
	}

	return m
}

// toModels returns an models.PetSlice
// this does nothing with the relationship templates
func (o PetTemplate) toModels(number int) models.PetSlice {
	m := make(models.PetSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.Pet
// according to the relationships in the template. Nothing is inserted into the db
func (t PetTemplate) setModelRels(o *models.Pet) {
	if t.r.Breed != nil {
		rel := t.r.Breed.o.toModel()
		rel.R.Pets = append(rel.R.Pets, o)
		o.BreedID = rel.ID
		o.R.Breed = rel
	}

	if t.r.OwnerUser != nil {
		rel := t.r.OwnerUser.o.toModel()
		rel.R.OwnerPets = append(rel.R.OwnerPets, o)
		o.OwnerID = rel.ID
		o.R.OwnerUser = rel
	}
}

// BuildSetter returns an *models.PetSetter
// this does nothing with the relationship templates
func (o PetTemplate) BuildSetter() *models.PetSetter {
	m := &models.PetSetter{}

	if o.ID != nil {
		m.ID = omit.From(o.ID())
	}
	if o.Name != nil {
		m.Name = omit.From(o.Name())
	}
	if o.Gender != nil {
		m.Gender = omit.From(o.Gender())
	}
	if o.Birthdate != nil {
		m.Birthdate = omit.From(o.Birthdate())
	}
	if o.BreedID != nil {
		m.BreedID = omit.From(o.BreedID())
	}
	if o.OwnerID != nil {
		m.OwnerID = omit.From(o.OwnerID())
	}

	return m
}

// BuildManySetter returns an []*models.PetSetter
// this does nothing with the relationship templates
func (o PetTemplate) BuildManySetter(number int) []*models.PetSetter {
	m := make([]*models.PetSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.Pet
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use PetTemplate.Create
func (o PetTemplate) Build() *models.Pet {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.PetSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use PetTemplate.CreateMany
func (o PetTemplate) BuildMany(number int) models.PetSlice {
	m := make(models.PetSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatablePet(m *models.PetSetter) {
	if m.Name.IsUnset() {
		m.Name = omit.From(random_string(nil))
	}
	if m.Birthdate.IsUnset() {
		m.Birthdate = omit.From(random_string(nil))
	}
	if m.BreedID.IsUnset() {
		m.BreedID = omit.From(random_uint32(nil))
	}
	if m.OwnerID.IsUnset() {
		m.OwnerID = omit.From(random_uint32(nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.Pet
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *PetTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.Pet) (context.Context, error) {
	var err error

	return ctx, err
}

// Create builds a pet and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *PetTemplate) Create(ctx context.Context, exec bob.Executor) (*models.Pet, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// MustCreate builds a pet and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// panics if an error occurs
func (o *PetTemplate) MustCreate(ctx context.Context, exec bob.Executor) *models.Pet {
	_, m, err := o.create(ctx, exec)
	if err != nil {
		panic(err)
	}
	return m
}

// CreateOrFail builds a pet and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// It calls `tb.Fatal(err)` on the test/benchmark if an error occurs
func (o *PetTemplate) CreateOrFail(ctx context.Context, tb testing.TB, exec bob.Executor) *models.Pet {
	tb.Helper()
	_, m, err := o.create(ctx, exec)
	if err != nil {
		tb.Fatal(err)
		return nil
	}
	return m
}

// create builds a pet and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *PetTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.Pet, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatablePet(opt)

	var rel0 *models.Breed
	if o.r.Breed == nil {
		var ok bool
		rel0, ok = breedCtx.Value(ctx)
		if !ok {
			PetMods.WithNewBreed().Apply(o)
		}
	}
	if o.r.Breed != nil {
		ctx, rel0, err = o.r.Breed.o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}
	opt.BreedID = omit.From(rel0.ID)

	var rel1 *models.User
	if o.r.OwnerUser == nil {
		var ok bool
		rel1, ok = userCtx.Value(ctx)
		if !ok {
			PetMods.WithNewOwnerUser().Apply(o)
		}
	}
	if o.r.OwnerUser != nil {
		ctx, rel1, err = o.r.OwnerUser.o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}
	opt.OwnerID = omit.From(rel1.ID)

	m, err := models.Pets.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = petCtx.WithValue(ctx, m)

	m.R.Breed = rel0
	m.R.OwnerUser = rel1

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple pets and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o PetTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.PetSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// MustCreateMany builds multiple pets and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// panics if an error occurs
func (o PetTemplate) MustCreateMany(ctx context.Context, exec bob.Executor, number int) models.PetSlice {
	_, m, err := o.createMany(ctx, exec, number)
	if err != nil {
		panic(err)
	}
	return m
}

// CreateManyOrFail builds multiple pets and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// It calls `tb.Fatal(err)` on the test/benchmark if an error occurs
func (o PetTemplate) CreateManyOrFail(ctx context.Context, tb testing.TB, exec bob.Executor, number int) models.PetSlice {
	tb.Helper()
	_, m, err := o.createMany(ctx, exec, number)
	if err != nil {
		tb.Fatal(err)
		return nil
	}
	return m
}

// createMany builds multiple pets and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o PetTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.PetSlice, error) {
	var err error
	m := make(models.PetSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// Pet has methods that act as mods for the PetTemplate
var PetMods petMods

type petMods struct{}

func (m petMods) RandomizeAllColumns(f *faker.Faker) PetMod {
	return PetModSlice{
		PetMods.RandomID(f),
		PetMods.RandomName(f),
		PetMods.RandomGender(f),
		PetMods.RandomBirthdate(f),
		PetMods.RandomBreedID(f),
		PetMods.RandomOwnerID(f),
	}
}

// Set the model columns to this value
func (m petMods) ID(val uint32) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.ID = func() uint32 { return val }
	})
}

// Set the Column from the function
func (m petMods) IDFunc(f func() uint32) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.ID = f
	})
}

// Clear any values for the column
func (m petMods) UnsetID() PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.ID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m petMods) RandomID(f *faker.Faker) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.ID = func() uint32 {
			return random_uint32(f)
		}
	})
}

// Set the model columns to this value
func (m petMods) Name(val string) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Name = func() string { return val }
	})
}

// Set the Column from the function
func (m petMods) NameFunc(f func() string) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Name = f
	})
}

// Clear any values for the column
func (m petMods) UnsetName() PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Name = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m petMods) RandomName(f *faker.Faker) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Name = func() string {
			return random_string(f)
		}
	})
}

// Set the model columns to this value
func (m petMods) Gender(val uint8) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Gender = func() uint8 { return val }
	})
}

// Set the Column from the function
func (m petMods) GenderFunc(f func() uint8) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Gender = f
	})
}

// Clear any values for the column
func (m petMods) UnsetGender() PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Gender = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m petMods) RandomGender(f *faker.Faker) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Gender = func() uint8 {
			return random_uint8(f)
		}
	})
}

// Set the model columns to this value
func (m petMods) Birthdate(val string) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Birthdate = func() string { return val }
	})
}

// Set the Column from the function
func (m petMods) BirthdateFunc(f func() string) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Birthdate = f
	})
}

// Clear any values for the column
func (m petMods) UnsetBirthdate() PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Birthdate = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m petMods) RandomBirthdate(f *faker.Faker) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.Birthdate = func() string {
			return random_string(f)
		}
	})
}

// Set the model columns to this value
func (m petMods) BreedID(val uint32) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.BreedID = func() uint32 { return val }
	})
}

// Set the Column from the function
func (m petMods) BreedIDFunc(f func() uint32) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.BreedID = f
	})
}

// Clear any values for the column
func (m petMods) UnsetBreedID() PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.BreedID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m petMods) RandomBreedID(f *faker.Faker) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.BreedID = func() uint32 {
			return random_uint32(f)
		}
	})
}

// Set the model columns to this value
func (m petMods) OwnerID(val uint32) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.OwnerID = func() uint32 { return val }
	})
}

// Set the Column from the function
func (m petMods) OwnerIDFunc(f func() uint32) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.OwnerID = f
	})
}

// Clear any values for the column
func (m petMods) UnsetOwnerID() PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.OwnerID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m petMods) RandomOwnerID(f *faker.Faker) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.OwnerID = func() uint32 {
			return random_uint32(f)
		}
	})
}

func (m petMods) WithBreed(rel *BreedTemplate) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.r.Breed = &petRBreedR{
			o: rel,
		}
	})
}

func (m petMods) WithNewBreed(mods ...BreedMod) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		related := o.f.NewBreed(mods...)

		m.WithBreed(related).Apply(o)
	})
}

func (m petMods) WithoutBreed() PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.r.Breed = nil
	})
}

func (m petMods) WithOwnerUser(rel *UserTemplate) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.r.OwnerUser = &petROwnerUserR{
			o: rel,
		}
	})
}

func (m petMods) WithNewOwnerUser(mods ...UserMod) PetMod {
	return PetModFunc(func(o *PetTemplate) {
		related := o.f.NewUser(mods...)

		m.WithOwnerUser(related).Apply(o)
	})
}

func (m petMods) WithoutOwnerUser() PetMod {
	return PetModFunc(func(o *PetTemplate) {
		o.r.OwnerUser = nil
	})
}
